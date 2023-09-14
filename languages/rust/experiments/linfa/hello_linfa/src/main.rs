use std::fs::File;
use std::io::Write;

use linfa::prelude::*;
use linfa_trees::DecisionTree;
use linfa_trees::SplitQuality;
use ndarray::prelude::*;
use ndarray::Array2;

fn main() {
    let original_data: Array2<f32> = array!(
        [1.0, 1.0, 1000.0, 1.0, 10.0],
        [1.0, 0.0, 0.0, 1.0, 6.0],
        [1.0, 0.0, 0.0, 1.0, 6.0],
        [1.0, 0.0, 0.0, 1.0, 6.0],
        [1.0, 0.0, 0.0, 1.0, 6.0],
        [1.0, 0.0, 800.0, 1.0, 8.0],
        [1.0, 0.0, 0.0, 0.0, 0.0],
        [1.0, 1.0, 0.0, 1.0, 9.0],
        [1.0, 1.0, 0.0, 1.0, 8.0],
        [1.0, 0.0, 800.0, 0.0, 8.0],
        [1.0, 1.0, 0.0, 1.0, 8.0],
        [1.0, 1.0, 500.0, 0.0, 8.0],
        [1.0, 0.0, 50.0, 0.0, 3.0],
        [1.0, 1.0, 50.0, 0.0, 4.0],
        [1.0, 0.0, 50.0, 0.0, 3.0],
    );

    let feature_names = vec!["Watched TV", "Pet Cat", "Rust LOC", "Ate Pizza"];
    let num_features = original_data.len_of(Axis(1)) - 1; // count everything except last column
    let features = original_data.slice(s![.., 0..num_features]).to_owned(); // get everything except the last column
    let labels = original_data.column(num_features).to_owned(); // get the last column

    let linfa_dataset = Dataset::new(features, labels)
        .map_targets(|x| match x.to_owned() as i32 {
            i32::MIN..=4 => "Sad",
            5..=7 => "Ok",
            _ => "Happy",
        })
        .with_feature_names(feature_names);

    let model = DecisionTree::params()
        .split_quality(SplitQuality::Gini)
        .fit(&linfa_dataset)
        .unwrap();

    File::create("output.tex")
        .unwrap()
        .write_all(model.export_to_tikz().with_legend().to_string().as_bytes())
        .unwrap();
}

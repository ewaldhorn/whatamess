extern crate globwalk;
use std::env;
use std::path::Path;
use std::process::Command;

use globwalk::DirEntry;

/**
 * Walks the folder structure and calls "cargo clean" on every Rust project.
 *
 * Useful for cleaning up some disk space in experimental project folders!
 */

fn main() -> Result<(), ()> {
    let base_path = env::current_dir().unwrap();

    let walker: Vec<DirEntry> =
        globwalk::GlobWalkerBuilder::from_patterns(base_path, &["Cargo.toml"])
            .max_depth(8)
            .follow_links(true)
            .build()
            .unwrap()
            .filter_map(Result::ok)
            .collect();

    println!("Walking {} folder(s).", walker.len());

    for dir in walker {
        let path = dir.path().parent().unwrap();
        let tmp = Path::new(path).join("target");
        let build_path = tmp.as_path();

        if build_path.exists() {
            let mut cmd = Command::new("cargo");
            cmd.arg("clean");
            cmd.current_dir(path);

            let cmd = cmd.output().unwrap();

            if !cmd.status.success() {
                println!("Failed to clean {:?}!", path);
                return Err(());
            } else {
                println!("Cleaned {:?}", path);
            }
        }
    }

    Ok(())
}

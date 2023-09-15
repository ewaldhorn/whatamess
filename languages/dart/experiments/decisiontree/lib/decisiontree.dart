import 'package:ml_algo/ml_algo.dart';
import 'package:ml_dataframe/ml_dataframe.dart';

Future<void> calculate() async {
  final samples = getDataFrame();

  // final pipeline = Pipeline(samples, [
  //   encodeAsIntegerLabels(
  //     featureNames: [
  //       'watched'
  //     ], // Here we convert strings from 'watched' column into numbers
  //   ),
  // ]);

  // pipeline.process(samples);

  final model = DecisionTreeClassifier(
    samples,
    'happiness',
    minError: 0.3,
    minSamplesCount: 5,
    maxDepth: 4,
  );

  await model.saveAsSvg('./file.svg');
}

encodeAsIntegerLabels({required List<String> featureNames}) {
  print('Received $featureNames');
  return 0;
}

DataFrame getDataFrame() {
  return DataFrame.fromRawCsv('''
watched tv,pet cat,rust loc,ate pizza,happiness
1,1,1000,1,10
1,0,0,1,6
1,0,0,1,6
1,0,0,1,6
1,0,0,1,6
1,0,800,1,8
1,0,0,0,0
1,1,0,1,9
1,1,0,1,8
1,0,800,0,8
1,1,0,1,8
1,1,500,0,8
1,0,50,0,3
1,1,50,0,4
1,0,50,0,3 
''');
//   return DataFrame.fromRawCsv('''
// watched tv,pet cat,rust loc,ate pizza,happiness
// yes,yes,1000,yes,10
// yes,no,0,yes,6
// yes,no,0,yes,6
// yes,no,0,yes,6
// yes,no,0,yes,6
// yes,no,800,yes,8
// yes,no,0,no,0
// yes,yes,0,yes,9
// yes,yes,0,yes,8
// yes,no,800,no,8
// yes,yes,0,yes,8
// yes,yes,500,no,8
// yes,no,50,no,3
// yes,yes,50,no,4
// yes,no,50,no,3
// ''');
}

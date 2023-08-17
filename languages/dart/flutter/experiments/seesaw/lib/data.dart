import 'package:charts_flutter/flutter.dart' as charts;

class SampleData {
  final String companySize;
  final int numberOfCompanies;

  SampleData(this.companySize, this.numberOfCompanies);

  static List<charts.Series<SampleData, String>> createChartData() {
    final data = [
      SampleData("1-15", 10),
      SampleData("15-50", 20),
      SampleData("51-200", 30),
      SampleData("201-500", 10),
      SampleData("501-1000", 40),
      SampleData("1000+", 50),
    ];

    return [
      charts.Series<SampleData, String>(
          id: 'SampleData',
          colorFn: (_, __) => charts.MaterialPalette.blue.shadeDefault,
          domainFn: (SampleData dataPoint, _) => dataPoint.companySize,
          measureFn: (SampleData dataPoint, _) => dataPoint.numberOfCompanies,
          data: data)
    ];
  }
}

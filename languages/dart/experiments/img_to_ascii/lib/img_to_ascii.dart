import 'dart:io';

import 'package:image/image.dart';

int calculate() {
  return 6 * 7;
}

void processImageToASCII(String inputfile, String outputfile) async {
  const width = 100;
  const bpp = 4;

  var pixels = grayscale(
          copyResize(decodeImage(new File(inputFile).readAsBytesSync()), width))
      .getBytes();
  var sink = File(outputFilename).openWrite();

  for (var i = 0; i < pixels.length; i += bpp) {
    if ((i / bpp) % width == 0) {
      sink.write("\n");
    }

    sink.write((pixels[i] + pixels[i + 1] + pixels[i + 2] > 0) ? '1' : '0');
  }

  await sink.flush();
  await sink.close();
}

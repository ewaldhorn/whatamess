import 'dart:io';

import 'package:image/image.dart';

void processImageToASCII(String inputFile, String outputFile) async {
  const width = 100;
  const bpp = 4;

  var tmpImage = decodeImage(File(inputFile).readAsBytesSync());

  if (tmpImage != null) {
    var pixels = grayscale(copyResize(tmpImage, width: width)).getBytes();
    var sink = File(outputFile).openWrite();

    for (var i = 0; i < pixels.length; i += bpp) {
      if ((i / bpp) % width == 0) {
        sink.write("\n");
      }

      sink.write((pixels[i] + pixels[i + 1] + pixels[i + 2] > 0) ? '1' : '0');
    }

    await sink.flush();
    await sink.close();
  } else {
    print('Unable to read the image, aborting.');
  }
}

import 'package:img_to_ascii/img_to_ascii.dart' as img_to_ascii;

void main(List<String> arguments) {
  print('Arguments: ${arguments.toString()}');
  if (arguments.length != 2) {
    print('Usage: app <img_source> <img_target>');
    return;
  }

  img_to_ascii.processImageToASCII(arguments.first, arguments.last);
}

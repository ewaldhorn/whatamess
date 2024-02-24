#include "hamming.h"

namespace hamming {
    int compute(string a, string b) {
        
        if (a.length() != b.length()) {
            throw std::domain_error("Strings are not of equal length!");
        }

        int count = 0;

        for (long unsigned int i = 0; i < a.length(); i++) {
            if (a[i] != b[i]) {
                count++;
            }
        }
        
        return count;
    }
}  // namespace hamming


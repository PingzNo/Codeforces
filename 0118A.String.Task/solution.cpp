#include <iostream>
#include <string>
#include <cctype>


void solution() {
    std::string text;
    std::cin >> text;

    for (std::string::const_iterator p = text.begin(); p != text.end(); ++p) {
        switch (tolower(*p)) {
            case 'a':
            case 'o':
            case 'y':
            case 'e':
            case 'u':
            case 'i':
                continue;

            default:
                std::cout << '.' << char(tolower(*p));
                break;
        }
    }

    std::cout << '\n';
}


void setup() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(NULL);
}


int main() {
    setup();
    solution();
}

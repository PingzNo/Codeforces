from typing import List


def solution():
    case_count = int(input())
    for case_index in range(case_count):
        text = input()
        pos = text.find('C')

        # R23C55 -> BC23
        if text[0] == 'R' and text[1].isdigit() and 1 < pos and pos < len(text):
            row = int(text[1:pos])
            col = int(text[pos + 1:])

            text_list: List[byte] = []
            while col > 0:
                if col % 26 == 0:
                    text_list.append('Z')
                    col = (col - 1) / 26
                else:
                    text_list.append('A'+(col % 26-1))
                    col /= 26

            while len(text_list) > 0:
                print(text_list[-1])
                text_list = text_list[:-1]
            print(row)

        else:
            # BC23 -> R23C55
            pos = 0
            while pos < len(text):
                if text[pos].isdigit():
                    break
                pos += 1

            row = text[:pos]
            col = text[pos:]

            c = 0
            for _, x in enumerate(row):
                c = c*26 + int(r-'A'+1)

            print("R{col}C{c}\n")


if __name__ == "__main__":
    solution()

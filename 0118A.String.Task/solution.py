def solution():
    for c in input().lower():
        if c == 'a' or c == 'o' or c == 'y' or c == 'e' or c == 'u' or c == 'i':
            continue

        print(f".{c}", end="")

    print()


if __name__ == "__main__":
    solution()

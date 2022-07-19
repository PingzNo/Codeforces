from typing import List

def solution():
    case_count = int(input())

    for _ in range(case_count):
        num_count = int(input())

        half_count = num_count // 2
        if half_count % 2 == 1:
            print("NO")
            continue

        half_half_count = half_count // 2
        ref_num = 0

        result: List[int] = []
        for num_index in range(half_half_count, 0, -1):
            result.append(ref_num - num_index * 2)

        for num_index in range(half_half_count):
            result.append(ref_num + num_index * 2)

        for num_index in range(half_half_count, 0, -1):
            result.append(ref_num - num_index * 2 - 1)

        for num_index in range(half_half_count):
            result.append(ref_num + num_index * 2 + 1)

        first = result[0]
        for num_index in range(len(result)):
            result[num_index] = result[num_index] - first + 2

        print("YES")
        print(" ".join([str(x) for x in result]))


if __name__ == "__main__":
	solution()

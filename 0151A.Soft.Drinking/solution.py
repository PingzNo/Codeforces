def solution():
	n, k, l, c, d, p, nl, np = [int(x) for x in input().split()]

	toast_count = k * l // nl
	slice_count = c * d
	salt_count = p // np

	result = min((toast_count, slice_count, salt_count)) // n
	print(result)


if __name__ == "__main__":
	solution()

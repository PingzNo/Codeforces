def solution():
	balance = int(input())

	if balance >= 0:
		print(balance)
		return

	balance = -balance
	if balance < 10:
		print("0")
		return

	last_digit = balance % 10
	digits = (balance - last_digit) // 10
	prev_digit = digits % 10

	if last_digit < prev_digit:
		result = -((digits - prev_digit) + last_digit)
	else:
		result = -((digits - prev_digit) + prev_digit)

	print(result)


if __name__ == "__main__":
	solution()

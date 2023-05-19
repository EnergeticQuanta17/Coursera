class Fraction:
	def __init__(self, a, b):
		self.a = a
		self.b = b

	def __str__(self):
		return f"({self.a}/{self.b})"

	def __mul__(self, f2):
		return Fraction(self.a * f2.a, self.b * f2.b)


f1 = Fraction(2,7)
f2 = Fraction(3,5)

f3 = f1 * f2
print(f3)
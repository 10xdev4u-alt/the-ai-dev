# Sample Python file for testing Mr.the-ai-dev
def fibonacci(n):
    """Generate fibonacci sequence up to n"""
    if n <= 1:
        return n
    else:
        return fibonacci(n-1) + fibonacci(n-2)

print("Hello from Prince's Python code!")
print(f"Fibonacci of 6: {fibonacci(6)}")
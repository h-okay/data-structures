class Matrix:
    def __init__(self, row_count: int, col_count: int, default: int = 0) -> "Matrix":
        self.data = [[default] * col_count for _ in range(row_count)]

    def __repr__(self):
        return f"Matrix({self.data})"

    def get(self, row: int, col: int) -> int | ValueError:
        if not self._exists(row, col):
            raise ValueError("index out of range")
        return self.data[row][col]

    def set(self, row: int, col: int, value: int) -> int | ValueError:
        if not self._exists(row, col):
            raise ValueError("index out of range")
        self.data[row][col] = value
        return value

    def _exists(self, row: int, col: int) -> bool:
        return 0 <= row < len(self.data) and 0 <= col < len(self.data[row])

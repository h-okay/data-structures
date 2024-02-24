import unittest
from array2d import Matrix


class TestMatrix(unittest.TestCase):
    def setUp(self):
        self.matrix = Matrix(3, 3)

    def test_get_set(self):
        self.matrix.set(1, 1, 5)
        value = self.matrix.get(1, 1)
        self.assertEqual(value, 5)

    def test_out_of_range(self):
        with self.assertRaises(ValueError):
            self.matrix.get(4, 4)
        with self.assertRaises(ValueError):
            self.matrix.set(4, 4, 10)

    def test_exists(self):
        self.assertTrue(self.matrix._exists(1, 1))
        self.assertFalse(self.matrix._exists(4, 4))

    def test_negative_indices(self):
        with self.assertRaises(ValueError):
            self.matrix.get(-1, 1)
        with self.assertRaises(ValueError):
            self.matrix.get(1, -1)
        with self.assertRaises(ValueError):
            self.matrix.set(-1, 1, 10)
        with self.assertRaises(ValueError):
            self.matrix.set(1, -1, 10)

    def test_empty_matrix(self):
        empty_matrix = Matrix(0, 0)
        with self.assertRaises(ValueError):
            empty_matrix.get(0, 0)
        with self.assertRaises(ValueError):
            empty_matrix.set(0, 0, 10)

    def test_out_of_range_indices(self):
        with self.assertRaises(ValueError):
            self.matrix.get(10, 10)
        with self.assertRaises(ValueError):
            self.matrix.set(10, 10, 10)


if __name__ == "__main__":
    unittest.main()

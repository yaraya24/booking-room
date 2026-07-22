#!/usr/bin/env python3
"""Capitalize all words in a text file.

Usage:
    python capitalize_words.py <input_file> [output_file]

If output_file is omitted, the input file is overwritten with the
capitalized content.
"""

import sys


def capitalize_file(input_path: str, output_path: str) -> None:
    with open(input_path, "r", encoding="utf-8") as f:
        content = f.read()

    capitalized = " ".join(word.capitalize() for word in content.split(" "))

    with open(output_path, "w", encoding="utf-8") as f:
        f.write(capitalized)


def main() -> None:
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} <input_file> [output_file]")
        sys.exit(1)

    input_path = sys.argv[1]
    output_path = sys.argv[2] if len(sys.argv) > 2 else input_path

    capitalize_file(input_path, output_path)
    print(f"Capitalized words written to {output_path}")


if __name__ == "__main__":
    main()

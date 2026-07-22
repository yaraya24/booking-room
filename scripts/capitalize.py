#!/usr/bin/env python3
"""Capitalize all letters in a text file.

Usage:
    python3 capitalize.py <input_file> [output_file]

If an output_file is provided, the capitalized text is written there.
Otherwise, the input file is overwritten in place.
"""

import argparse
import sys


def capitalize_file(input_path: str, output_path: str) -> None:
    with open(input_path, "r", encoding="utf-8") as f:
        content = f.read()

    with open(output_path, "w", encoding="utf-8") as f:
        f.write(content.upper())


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Capitalize all letters in a text file."
    )
    parser.add_argument("input_file", help="Path to the input text file")
    parser.add_argument(
        "output_file",
        nargs="?",
        default=None,
        help="Path to write the capitalized text (defaults to overwriting input_file)",
    )
    args = parser.parse_args()

    output_path = args.output_file or args.input_file

    try:
        capitalize_file(args.input_file, output_path)
    except FileNotFoundError:
        print(f"Error: file not found: {args.input_file}", file=sys.stderr)
        sys.exit(1)

    print(f"Capitalized text written to {output_path}")


if __name__ == "__main__":
    main()

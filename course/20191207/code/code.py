#encoding: utf-8
import os

if __name__ == '__main__':
    total = 0
    for path, _, names in os.walk("."):
        for name in names:
            if not name.endswith(".go"):
                continue
            lines = open(os.path.join(path, name), mode="rt", encoding="utf-8").readlines()
            total += len([line for line in lines if line.strip() != ""])

    print("line: ", total)
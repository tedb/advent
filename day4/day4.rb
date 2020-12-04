puts "Part 1: ", File.read("day4/input").split("\n\n").count { |p|
    colons = p.count(":")
    colons == 8 || (colons == 7 && ! p.match(/cid:/))
}

pattern = /
(?:
    # byr (Birth Year) - four digits; at least 1920 and at most 2002.
    (byr:(?:19[2-8][0-9]|199[0-9]|200[0-2]))
    (?:\s+|$)|

    # iyr (Issue Year) - four digits; at least 2010 and at most 2020.
    (iyr:(?:201[0-9]|2020))
    (?:\s+|$)|

    # eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
    (eyr:(?:202[0-9]|2030))
    (?:\s+|$)|

    # hgt (Height) - a number followed by either cm or in:
    # If cm, the number must be at least 150 and at most 193.
    # If in, the number must be at least 59 and at most 76.
    (hgt:(?:(?:1[5-8][0-9]|19[0-3])cm|(?:59|6[0-9]|7[0-6])in))
    (?:\s+|$)|

    # hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
    (hcl:\#[0-9a-f]{6})
    (?:\s+|$)|

    # ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    (ecl:(?:amb|blu|brn|gry|grn|hzl|oth))
    (?:\s+|$)|

    # pid (Passport ID) - a nine-digit number, including leading zeroes.
    (pid:[0-9]{9})
    (?:\s+|$)|

    # cid (Country ID) - ignored, missing or not.
    \S*
    (?:\s+|$)
)+
/x

puts "Part 2: ", File.read("day4/input").split("\n\n").count { |p| p.match(pattern).captures.reject(&:nil?).count == 7 }

import z3
import functools
import sys

lines = [line.strip() for line in sys.stdin]

# Helpers for 64 bit integer arithmetic
Int64Val = lambda val: z3.BitVecVal(val, 64)
Int64 = lambda name: z3.BitVec(name, 64)

# Initialize machine registers
regnames = "wxyz"
registers = {reg : Int64Val(0) for reg in regnames}
get_value = lambda s: registers[s] if s in regnames else Int64Val(int(s))

# Inputs and bounds
inputs = [Int64(f"input{i}") for i in range(14)]
input_bounds = [z3.And(i > Int64Val(0), i < Int64Val(10)) for i in inputs]
input_bounds = z3.simplify(functools.reduce(z3.And, input_bounds, True))
input_counter = 0

for line in lines:
    tokens = line.split(" ")
    cmd, reg = tokens[0], tokens[1]
    arg = tokens[2] if len(tokens) > 2 else None
    if cmd == "inp":
        registers[reg] = inputs[input_counter]
        input_counter += 1
    elif cmd == "add":
        val = get_value(arg)
        registers[reg] += val
    elif cmd == "mul":
        val = get_value(arg)
        registers[reg] *= val
    elif cmd == "div":
        val = get_value(arg)
        registers[reg] /= val
    elif cmd == "mod":
        val = get_value(arg)
        registers[reg] %= val
    elif cmd == "eql":
        val = get_value(arg)
        registers[reg] = z3.If(registers[reg] == val, 
            Int64Val(1), 
            Int64Val(0)
        )
    else:
        print("Unexpected line:", line)
        exit(1)

s = z3.Optimize()

def get_solution(solver):
    if solver.check() == z3.sat:
        model = solver.model()
        solution = ""
        for input_var in inputs:
            solution += str(model[input_var].as_long())
        return solution
    return None

constraint = registers['z'] == Int64Val(0)
s.add(constraint)
s.add(input_bounds)

# Encode the 14-digit value represented by input variables
objective = 0
for (p, i) in enumerate(inputs[::-1]):
    objective = objective + Int64Val(10**p) * i

s.push()
s.maximize(objective)
print("Maximizing")
print(get_solution(s))

s.pop()
s.minimize(objective)
print("Minimizing")
print(get_solution(s))
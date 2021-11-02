x := linspace(0, 2*pi, 20)
y := [sin(x) for i <- x]
xlabel("x")
ylabel("sin(x)")
plot(x, y)
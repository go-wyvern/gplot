x := linspace(0, 2*pi, 20)
y := [sin(i) for i <- x]
xlabel("x")
ylabel("sin(x)")
plot(x, y)
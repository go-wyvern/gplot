x := linspace(0, 2*pi, 20)
y1 := [sin(x) for i <- x]
y2 := [cos(x) for i <- x]
y3 := [x*x for i <- x]
y4 := [-1*x*x for i <- x]
subplot(2,2,1)
plot(x, y1)
subplot(2,2,2)
plot(x, y2)
subplot(2,2,3)
plot(x, y3)
subplot(2,2,4)
plot(x, y4)
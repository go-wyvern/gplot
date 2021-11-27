x := linspace(0, 2*pi, 20)
y1 := [sin(i) for i <- x]
y2 := [cos(i) for i <- x]
y3 := [i*i for i <- x]
y4 := [-1*i*i for i <- x]
subplot(2,2,1)
plot(x, y1)
subplot(2,2,2)
plot(x, y2)
subplot(2,2,3)
plot(x, y3)
subplot(2,2,4)
plot(x, y4)
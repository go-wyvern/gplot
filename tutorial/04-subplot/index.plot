x := linspace(0, 2*pi, 20)
subplot(2,2,1)
plot(x, sin(x))
subplot(2,2,2)
plot(x, cos(x))
subplot(2,2,3)
plot(x, x*x)
subplot(2,2,4)
plot(x, scale(2,x)*x)

prod_str = 'data/producer_';
con_str = 'data/client_';
txt = '.txt';
con_avgs = zeros(1, 20);
prod_avgs = zeros(1, 20);
for i = 1:20
    con_avgs(i) = mean(load(strcat(con_str, num2str(i * 10), txt))(1:10));
    prod_avgs(i) = mean(load(strcat(prod_str, num2str(i * 10), txt))(1:10));
endfor

figure;
hold on;
plot(10:10:200, con_avgs, 'b+-', 'linewidth', 3)
plot(10:10:200, prod_avgs, 'r+-', 'linewidth', 3)
title('Time taken per request')
xlabel('Number of producers / consumers')
ylabel('Time (seconds)')
legend('Consumer', 'Producer')
print('figs/clpr.png', '-dpng');

waitforbuttonpress


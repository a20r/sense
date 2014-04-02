
prod_str = 'data/producer_';
con_str = 'data/client_';
txt = '.txt';
con_avgs = zeros(1, 20);
prod_avgs = zeros(1, 20);
prod_req = zeros(1, 20);
con_req = zeros(1, 20);
for i = 1:20
    con_data = load(strcat(con_str, num2str(i * 10), txt));
    prod_data = load(strcat(prod_str, num2str(i * 10), txt));
    con_avgs(i) = mean(con_data);
    prod_avgs(i) = mean(prod_data);
    prod_req(i) = size(prod_data)(1) / (i);
    con_req(i) = size(con_data)(1) / (i);
endfor

prod_req(find(prod_req > 100)) = 100
con_req(find(con_req > 100)) = 100

figure;
hold on;
plot(10:10:200, con_avgs, 'b+-', 'linewidth', 3)
plot(10:10:200, prod_avgs, 'r+-', 'linewidth', 3)
title('Time taken per request')
xlabel('Number of producers / consumers')
ylabel('Time (seconds)')
legend('Consumer', 'Producer')
print('figs/clpr.png', '-dpng');

hold off;
figure;
plot(10:10:200, con_req, 'b+-', 'linewidth', 3)
hold on;
plot(10:10:200, prod_req, 'r+-', 'linewidth', 3)
title('Request reliability')
xlabel('Number of producers / consumers')
ylabel('% of data transfer')
legend('Consumer', 'Producer')
ylim([75 120])
print('figs/reqs.png', '-dpng');


waitforbuttonpress


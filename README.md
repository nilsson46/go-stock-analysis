Scraping could be a option... 

Start to add some manually to the db or in the frontend later on. 


simon@Simons-MBP go-stock-analysis % docker-compose up -d
Creating network "go-stock-analysis_default" with the default driver
Creating go-stock-analysis_db_1 ... done
Creating go-stock-analysis_app_1 ... done
simon@Simons-MBP go-stock-analysis % docker-compose logs app
Attaching to go-stock-analysis_app_1
app_1  | 2024/11/21 20:37:20 Unable to connect to database: failed to connect to `host=localhost user=postgres database=stock_analysis_db`: dial error (dial tcp [::1]:5432: connect: connection refused)
simon@Simons-MBP go-stock-analysis %  
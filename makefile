deploy:
	make -j 2 deploy_chart deploy_options

deploy_chart:
	gcloud functions deploy chart-data --source="src/" --entry-point="ChartData" --allow-unauthenticated

deploy_options:
	gcloud functions deploy options --source="src/" --entry-point="Options" --allow-unauthenticated

serve:
	make -j 2 serve_options serve_chart

serve_chart:
	go build -o server/chart-data/main.exe github.com/bryanplant/covid-charts/server/chart-data
	server/chart-data/main.exe

serve_options:
	go build -o server/options/main.exe github.com/bryanplant/covid-charts/server/options
	server/options/main.exe

auth:
	gcloud auth print-identity-token
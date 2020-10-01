deploy:
	make -j 2 deploy_chart deploy_options

deploy_chart:
	gcloud functions deploy chart-data --source="src/" --entry-point="ChartData" --allow-unauthenticated

deploy_options:
	gcloud functions deploy options --source="src/" --entry-point="Options" --allow-unauthenticated

serve:
	set FIRESTORE_EMULATOR_HOST=localhost:8081
	go build -o server/main.exe github.com/bryanplant/covid-charts/server
	server/main.exe

auth:
	gcloud auth print-identity-token
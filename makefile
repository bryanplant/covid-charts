deploy:
	make -j 3 deploy_chart deploy_options deploy_update

deploy_chart:
	gcloud functions deploy chart-data --entry-point="ChartData" --allow-unauthenticated

deploy_options:
	gcloud functions deploy options --entry-point="Options" --allow-unauthenticated

deploy_update:
	gcloud functions deploy update-data --entry-point="UpdateData" --allow-unauthenticated

serve:
	set FIRESTORE_EMULATOR_HOST=localhost:8081
	go build -o server/main.exe github.com/bryanplant/covid-charts/server
	server/main.exe

firestore:
	gcloud beta emulators firestore start --host-port=localhost:8081

auth:
	gcloud auth print-identity-token
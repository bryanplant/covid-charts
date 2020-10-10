deploy:
	make -j 5 deploy_chart deploy_options deploy_update deploy_save deploy load

deploy_chart:
	gcloud functions deploy chart-data --entry-point="ChartData" --allow-unauthenticated --trigger-http --runtime=go113

deploy_options:
	gcloud functions deploy options --entry-point="Options" --allow-unauthenticated --trigger-http --runtime=go113

deploy_update:
	gcloud functions deploy update-data --entry-point="UpdateData" --allow-unauthenticated --trigger-http --runtime=go113

deploy_save:
	gcloud functions deploy save --entry-point="SaveChart" --allow-unauthenticated --trigger-http --runtime=go113

deploy_load:
	gcloud functions deploy load --entry-point="LoadChart" --allow-unauthenticated --trigger-http --runtime=go113

serve:
	set FIRESTORE_EMULATOR_HOST=localhost:8081
	go build -o server/main.exe github.com/bryanplant/covid-charts/server
	server/main.exe

firestore:
	gcloud beta emulators firestore start --host-port=localhost:8081

auth:
	gcloud auth print-identity-token
docker build -f Dockerfile -t gig-service-4-scripts .
docker run --rm -p 9000:9000 -d --name local_gig_service_4_scripts gig-service-4-scripts

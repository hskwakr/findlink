FROM httpd:buster

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		wget \
		build-essential

# Install go
RUN wget -O ./go.tar.gz https://golang.org/dl/go1.17.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz

# Copy test page
COPY ./*.html /usr/local/apache2/htdocs/

WORKDIR /home/working

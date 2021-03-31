FROM ubuntu

RUN wget -q -O - https://packages.cloudfoundry.org/debian/cli.cloudfoundry.org.key | sudo apt-key add - \
    && echo "deb https://packages.cloudfoundry.org/debian stable main" | sudo tee /etc/apt/sources.list.d/cloudfoundry-cli.list\
    && apt-get update \
    && apt-get install -y --no-install-recommends cf-cli \
    && rm -rf /var/lib/apt/lists/*

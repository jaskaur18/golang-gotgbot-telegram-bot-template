#!/bin/bash

# Function to check if a package is installed
is_package_installed() {
  if command -v "$1" >/dev/null 2>&1; then
    return 0
  else
    return 1
  fi
}

# Function to install a package if not already installed
install_package() {
  if ! is_package_installed "$1"; then
    echo "Installing $1..."
    sudo apt update
    sudo apt install -y "$1"
  else
    echo "$1 is already installed."
  fi
}

# Function to setup environment variables
setup_env_variables() {
  if ! grep -q "$1" ~/.bashrc; then
    echo "Setting up $1 environment variable..."
    echo "export $1=$2" >>~/.bashrc
    source ~/.bashrc
  else
    echo "$1 environment variable is already set up."
  fi
}

# Function to configure PostgreSQL for remote access
configure_postgresql() {
  sudo sed -i "s/#listen_addresses = 'localhost'/listen_addresses = '*'/" /etc/postgresql/*/main/postgresql.conf
  echo "host    all             all             0.0.0.0/0               md5" | sudo tee -a /etc/postgresql/*/main/pg_hba.conf >/dev/null
  sudo -u postgres psql -c "CREATE USER $1 WITH ENCRYPTED PASSWORD '$2';"
  sudo -u postgres psql -c "ALTER USER $1 CREATEDB CREATEROLE;"
}

# Function to install Node.js using fnm
install_nodejs() {

  if ! fnm list-remote | grep -q "$1"; then
    echo "Installing Node.js $1 using fnm..."
    fnm install "$1"
    fnm use "$1"
  else
    echo "Node.js version $1 is already installed."
  fi
}

# Function to install a global npm package
install_global_npm_package() {
  if ! is_package_installed "$1"; then
    echo "Installing $1 using npm..."
    npm install -g "$1"
  else
    echo "$1 is already installed."
  fi
}

# Main script
echo "Starting setup..."

#Update and upgrade
sudo apt update
sudo apt upgrade -y

install_package "wget"
install_package "unzip"
install_package "make"
install_package "clang"

if ! install_package "wget"; then
  echo "Error installing wget"
  exit 1
fi

if ! install_package "unzip"; then
  echo "Error installing unzip"
  exit 1
fi

setup_env_variables "GOPATH" "\$HOME/go"
setup_env_variables "PATH" "\$PATH:\$GOPATH/bin"

if ! is_package_installed "golang"; then
  echo "Downloading and installing Golang..."
  wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
  tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
  rm go1.21.0.linux-amd64.tar.gz
  echo 'export PATH=$PATH:/usr/local/go/bin' >>~/.bashrc
  source ~/.bashrc
else
  echo "Golang is already installed."
fi

install_package "postgresql"

read -p "Enter PostgreSQL username: " pg_username
read -s -p "Enter password for $pg_username: " pg_password

if ! configure_postgresql "$pg_username" "$pg_password"; then
  echo "Error configuring PostgreSQL"
  exit 1
fi

bash -c "$(curl -fsSL https://fnm.vercel.app/install)"

echo 'export PATH="/root/.local/share/fnm:$PATH"' >>~/.bashrc
echo 'eval "`fnm env`"'
source ~/.bashrc

install_nodejs "20"

install_global_npm_package "pm2"

echo "Setup complete!"

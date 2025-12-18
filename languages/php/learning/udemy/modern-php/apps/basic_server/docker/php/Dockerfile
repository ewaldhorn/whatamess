# Use the official PHP 8.4 FPM image as the base
FROM php:8.4-fpm-alpine AS dev

# Install core PHP extensions (e.g., mysqli, pdo_mysql, and opcache)
RUN apk update && apk add --no-cache \
    autoconf \
    g++ \
    make \
    # Packages needed for PHP extensions
    libzip-dev \
    libpng-dev \
    icu-dev \
    # Install common extensions
    && docker-php-ext-install pdo_mysql opcache intl zip gd \
    # Clean up
    && rm -rf /var/cache/apk/*

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Set the working directory for the application
WORKDIR /var/www/html

# Expose the PHP-FPM port (internal, Nginx connects to this)
EXPOSE 9000

CMD ["php-fpm"]

<?php
// app/index.php
echo "<h1>Hello from PHP " . phpversion() . "!</h1>";

// Test MySQL connection (replace credentials with those in docker-compose.yml)
$host = 'db'; // This is the service name from docker-compose!
$db   = 'my_api_db';
$user = 'api_user';
$pass = 'api_user_password';
$charset = 'utf8mb4';

$dsn = "mysql:host=$host;dbname=$db;charset=$charset";
$options = [
    PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION,
    PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
    PDO::ATTR_EMULATE_PREPARES   => false,
];

try {
     $pdo = new PDO($dsn, $user, $pass, $options);
     echo "<p style='color: green;'>✅ Database connection successful!</p>";
} catch (\PDOException $e) {
     echo "<p style='color: red;'>❌ Database connection failed: " . $e->getMessage() . "</p>";
}
?>

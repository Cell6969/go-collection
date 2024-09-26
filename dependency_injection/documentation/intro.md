# Intro

Pada umumnya ketika membuat aplikasi khusus di backend sering mengenal dengan konsep *Dependency Injection* (DI). DI merupakan konsep dimana sebuah object membutuhkan object lain. Seperti contoh misal controller membutuhkan service dan service membutuhkan repository.

Untuk membuat dependency injection biasanya membutuhkan constructor, namun karena di golang tidak berbasis object oriented maka constructor dilakukan dengan menggunakan function.

Untuk melakukan dependency injection secara cepat biasanya menggunakan library :

1. google wire
2. uber-go fx
3. golobby container.

Untuk sesi ini akan menggunakan google wire. Lalu untuk base menggunakan project dari restful api
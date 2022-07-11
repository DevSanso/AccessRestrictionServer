plugins {
    id("java")
}

group = "com.github.DevSanso.accessRestrictionGateway"
version = "unspecified"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.springframework.boot:spring-boot-starter-test:2.7.1")
    testImplementation("junit:junit:4.13.2")
    implementation("org.springframework.boot:spring-boot-starter-data-jpa:2.7.1")
    testImplementation("com.h2database:h2:2.1.214")
    compileOnly("org.projectlombok:lombok:1.18.24")
    annotationProcessor("org.projectlombok:lombok:1.18.24")
}


tasks.getByName<Test>("test") {
    useJUnitPlatform()
}
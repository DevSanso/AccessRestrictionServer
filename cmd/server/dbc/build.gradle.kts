plugins {
    id("java")
}

group = "com.github.DevSanso.accessRestrictionGateway"
version = "unspecified"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.8.1")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.8.1")
    implementation("org.springframework.boot:spring-boot-starter-data-jpa:2.7.1")
    implementation("org.xerial:sqlite-jdbc:3.36.0.3")
}


tasks.getByName<Test>("test") {
    useJUnitPlatform()
}
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
}

java

tasks.getByName<Test>("test") {
    useJUnitPlatform()
}
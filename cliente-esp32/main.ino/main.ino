#include <WiFi.h>
#include <HTTPClient.h>

// --- DADOS GLOBAIS ---
// Informação que descobrimos: o pino do LED!
const int LED_PIN = 4;

// --- CONFIGURE AQUI ---
const char* ssid = "IGOR_2.4";
const char* password = "66999983273";
// Lembre-se de verificar o IP do seu PC com 'ipconfig' ou 'ip addr'
String serverUrl = "http://192.168.0.8:8080/stats"; // MUDE PARA O IP DO SEU PC
// --------------------

void setup() {
  Serial.begin(115200); // Inicia a comunicação serial
  pinMode(LED_PIN, OUTPUT); // Configura o pino do LED como saída
  delay(1000);

  Serial.print("Conectando ao WiFi: ");
  Serial.println(ssid);

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("\nWiFi conectado!");
  Serial.print("Endereço IP do ESP32: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  if (WiFi.status() == WL_CONNECTED) {
    // Pisca o LED para indicar que uma nova requisição vai começar
    digitalWrite(LED_PIN, HIGH);
    delay(100);
    digitalWrite(LED_PIN, LOW);
    
    HTTPClient http;
    Serial.println("\n[HTTP] Iniciando requisição para: " + serverUrl);
    http.begin(serverUrl);

    int httpCode = http.GET();

    if (httpCode > 0) {
      Serial.printf("[HTTP] Código de retorno: %d\n", httpCode);
      if (httpCode == HTTP_CODE_OK) {
        String payload = http.getString();
        Serial.println("Resposta recebida:");
        Serial.println(payload);
      }
    } else {
      Serial.printf("[HTTP] Falha na requisição, erro: %s\n", http.errorToString(httpCode).c_str());
    }
    http.end();
  } else {
    Serial.println("WiFi desconectado.");
  }
  delay(5000); // Espera 5 segundos
}
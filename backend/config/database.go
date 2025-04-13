package config

import (
    "context"
    "log"
    "sync"

    "github.com/spf13/viper"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    clientOnce     sync.Once
)

func init() {
    // Configuração do Viper
    viper.SetConfigName("config") // Nome do arquivo (sem extensão)
    viper.SetConfigType("yaml")   // Tipo do arquivo
    viper.AddConfigPath("./config") // Caminho para o arquivo de configuração
    viper.AutomaticEnv()           // Carrega variáveis de ambiente automaticamente

    // Leia o arquivo de configuração
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
    }
}

// GetDatabase retorna uma instância do banco de dados MongoDB
func GetDatabase() *mongo.Database {
    clientOnce.Do(func() {
        mongoURI := viper.GetString("mongo.uri")
        if mongoURI == "" {
            log.Fatal("A configuração 'mongo.uri' não foi encontrada")
        }

        clientOptions := options.Client().ApplyURI(mongoURI)
        client, err := mongo.Connect(context.Background(), clientOptions)
        if err != nil {
            log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
        }
        clientInstance = client
    })

    dbName := viper.GetString("mongo.db_name")
    if dbName == "" {
        log.Fatal("A configuração 'mongo.db_name' não foi encontrada")
    }

    return clientInstance.Database(dbName)
}
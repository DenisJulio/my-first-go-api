CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    language CHAR(5) NOT NULL
);

INSERT INTO
    messages (message, language)
VALUES
    ('Hello World!', 'en-US'),
    ('Bonjour le monde!', 'fr-FR'),
    ('Hola Mundo!', 'es-ES'),
    ('Ciao Mondo!', 'it-IT'),
    ('Hallo Welt!', 'de-DE'),
    ('Olá Mundo!', 'pt-BR'),
    ('Привет, мир!', 'ru-RU');
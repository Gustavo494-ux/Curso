INSERT INTO usuarios(nome,nick,email,senha) VALUES
    ("Usuário 1 ","usuario_1","usuario1@gmail.com","$2a$10$nF0nF3Ti1.gVBpBJO4SVLeqN7Rba1/9arPtLzMN9yAHn9obub4pza"),
    ("Usuário 2 ","usuario_2","usuario2@gmail.com","$2a$10$nF0nF3Ti1.gVBpBJO4SVLeqN7Rba1/9arPtLzMN9yAHn9obub4pza"),
    ("Usuário 3 ","usuario_3","usuario3@gmail.com","$2a$10$nF0nF3Ti1.gVBpBJO4SVLeqN7Rba1/9arPtLzMN9yAHn9obub4pza");

INSERT INTO seguidores(usuario_id, seguidor_id) VALUES
    (1,2),
    (3,1),
    (1,3);

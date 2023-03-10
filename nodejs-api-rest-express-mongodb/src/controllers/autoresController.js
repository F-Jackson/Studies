import autores from "../models/Autor.js";

export default class AutorController {
    static listarAutores = (req, res) => {
        autores.find((err, autores) => {
            res.status(200).send(autores)
        })
    }

    static listarAutorPorId = (req, res) => {
        const {id} = req.params;

        autores.findById(id, (err, autores) => {
            if(err) {
                res.status(400).send({message:`${err.message} - id  do autor nao localizado`})
            } else {
                res.status(200).send(autores);
            }
        })
    }

    static cadastrarAutor = (req, res) => {
        let autor = new autores(req.body);

        autor.save((err) => {
            if(err) {
                res.status(500).send({message: `${err.message} - falha ao cadatrar autor`});
            } else {
                res.status(201).send(autor.toJSON());
            }
        })
    }

    static atualizarAutor = (req, res) => {
        const {id} = req.params;
        autores.findByIdAndUpdate(id, {$set: req.body}, (err)=> {
            if(!err) {
                res.status(200).send({message: "autor atualizado com sucesso"})
            } else {
                res.status(500).send({message: err.message });
            }
        });
    }

    static excluirAutor = (req, res) => {
        const {id} = req.params;
        autores.findByIdAndDelete(id, (err) => {
            if(!err) {
                res.status(200).send({message: 'autor excluido com sucesso'})
            } else {
                res.status(500).send({message: err.message})
            }
        });
    }
}
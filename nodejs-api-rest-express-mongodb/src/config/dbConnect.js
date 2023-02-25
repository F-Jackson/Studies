import mongoose from "mongoose";

mongoose.connect("mongodb+srv://fjacksondev:851Wetq3HiSyzSeM@cluster0.e9fucxi.mongodb.net/Alura-Node");

let db = mongoose.connection;

export default db;
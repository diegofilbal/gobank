import axios from "axios";
import React, { useRef, useState } from "react";
import { toast } from "react-toastify";
import styled from "styled-components";

const FormContainer = styled.form`
    display: flex;
    align-items: flex-end;
    gap: 10px;
    flex-wrap: wrap;
    background-color: #fff;
    padding: 20px;
    box-shadow: 0px 0px 5px #ccc;
    border-radius: 5px;
`;

const InputArea = styled.div`
    display: flex;
    flex-direction:column;
`;

const Input = styled.input`
    width:120px;
    padding:0 10px;
    border: 1px solid #bbb;
    border-radius: 5px;
    height: 40px;
`;

const Label = styled.label``;

const Button = styled.button`
padding: 10px;
cursor: pointer;
border-radius: 5px;
border: none;
background-color: #2c73d2;
color: white;
height: 42px;
`;

const Form = () => {
    const ref = useRef();
    const [numero, setNumero] = useState(Number);
    const [tipo, setTipo] = useState(String);
    const [saldo, setSaldo] = useState(Number);
    
    const handleSubmit = async (e:React.FormEvent<HTMLFormElement>)=>{
        e.preventDefault(); 
        await axios.post("http://localhost:8001/banco/conta",{
            "numero" : numero, 
            "tipo" : tipo,
            "saldoInicial" : saldo
        })
        .then(()=>{ toast.success("Reserva feita com sucesso")})
        .catch(({data})=>toast.error(data));
        
    };

    return (
        <FormContainer onSubmit={handleSubmit}>
            <InputArea>
                <Label>Numero da conta</Label>
                <Input name="qtd" type= "number" onChange={(e)=>setNumero(parseInt(e.target.value))}></Input>
            </InputArea>
            <InputArea>
                <Label>Normal</Label>
                <Input name="tipo" type="radio" onChange={(e)=>setTipo("Normal")}></Input>
                </InputArea>
            <InputArea>
                <Label>Bônus</Label>
                <Input name="tipo" type="radio" onChange={(e)=>setTipo("Bônus")}></Input>
                </InputArea>
            <InputArea>
                <Label>Poupanca</Label>
                <Input name="tipo" type="radio" onChange={(e)=>setTipo("Poupança")}></Input>
            </InputArea>
            <InputArea>
                <Label>Saldo Inicial</Label>
                <Input name="saldo" type= "number" onChange={(e)=>setSaldo(parseInt(e.target.value))}></Input>
            </InputArea>
            <Button type="submit">Enviar</Button>
        </FormContainer>
    );
};

export default Form;
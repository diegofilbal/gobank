import React, { useEffect, useRef, useState } from "react";
import styled from "styled-components";
import axios from "axios";
import { toast } from "react-toastify";
import {FaTrash, FaEdit, FaCheck, FaCalendarCheck, FaCalendarAlt, FaCalendarTimes} from "react-icons/fa";
import { resolve4 } from "dns";

const Table = styled.table`
    width: 100%;
    background-color: #fff;
    padding: 20px;
    box-shadow: 0px 0px 5px #ccc;
    border-radius: 5px;
    margin: 20px auto;
    word-break: break-all;
`;

export const Thead = styled.thead``

export const Tr = styled.tr``;

export const Td = styled.th`
    text-align: center;
    width: auto;
    padding: 10px;   
    padding-top: 15px; 
`;

export const Th = styled.th`
    text-align: start;
    border-bottom: inset;
    padding-bottom: 5px;    
    padding-left: 5px;
`;

export interface reservasType {
    id: Number,
	data_reserva: String,
	qtd_pessoas: Number,
	checkin_previsto: String,
    checkin: null | String,
    checkout_previsto: String,
    checkout: null | String,
    nome: String,
    andar: String
};

const Button = styled.button`
padding: 10px;
cursor: pointer;
border-radius: 5px;
border: none;
background-color: #2c73d2;
color: white;
height: 42px;
`;


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

type ContaType ={
    pontuacao: Number;
    saldo: Number;
    tipo: String;
};

const Grid = () => {
    const ref = useRef();
    const [contaNum, setContaNum] = useState(Number) ;
    const [valorCredito, setValorCredito] = useState(Number) ;
    const [valorDebito, setValorDebito] = useState(Number) ;
    
    const [taxaRendimento, setTaxaRendimento] = useState(Number) ;
    
    const [numContaCredito, setNumContaCredito] = useState(Number) ;
    const [numContaDebito, setNumContaDebito] = useState(Number) ;
    const [valorTransferencia, setValorTransferencia] = useState(Number) ;
    const [conta, setConta] = useState<ContaType>({
        pontuacao: 0,
        saldo: 0,
        tipo: "Normal"

    });
  
  const getConta = async(e:React.FormEvent<HTMLFormElement>) =>{
    e.preventDefault();     
    try{
       await axios.get(`http://localhost:8001/banco/conta/${contaNum}`)
      .then((response)=>{setConta(response.data); toast.success("conta Encontrada!")})
      .catch((error)=>{
          if (!(!error.response)) {
              // A requisição foi feita e o servidor respondeu com um código de status
              // que sai do alcance de 2xx
              toast.error(error.response.status + " " + error.response.data.error);
            } else if (error.request) {
              // A requisição foi feita mas nenhuma resposta foi recebida
              // `error.request` é uma instância do XMLHttpRequest no navegador e uma instância de
              // http.ClientRequest no node.js
              toast.error(error.request);
            } else {
              // Alguma coisa acontenceu ao configurar a requisição que acionou este erro.
              toast.error('Error', error.message);
            }
            toast.error(error.config);
      });
    }catch( error : any){
      toast.error( error);
    }
  }

  const postCredito = async(e:React.FormEvent<HTMLFormElement>) =>{
    e.preventDefault();     
    try{
      const res = await axios.put(`http://localhost:8001/banco/conta/${contaNum}/credito`,{
        "valor": valorCredito,
      })
      .then((response)=>{ toast.success(response.data.message)})
      .catch((error)=>{
          if (!(!error.response)) {
              // A requisição foi feita e o servidor respondeu com um código de status
              // que sai do alcance de 2xx
              toast.error(error.response.status + " " + error.response.data.error);
            } else if (error.request) {
              // A requisição foi feita mas nenhuma resposta foi recebida
              // `error.request` é uma instância do XMLHttpRequest no navegador e uma instância de
              // http.ClientRequest no node.js
              toast.error(error.request);
            } else {
              // Alguma coisa acontenceu ao configurar a requisição que acionou este erro.
              toast.error('Error', error.message);
            }
            toast.error(error.config);
      });
    }catch( error : any){
      toast.error( error);
    }
  }

  const postDebito = async(e:React.FormEvent<HTMLFormElement>) =>{
    e.preventDefault();     
    try{
      const res = await axios.put(`http://localhost:8001/banco/conta/${contaNum}/debito`,{
        "valor": valorDebito,
      })
      .then((response)=>{toast.success(response.data.message)})
      .catch((error)=>{
          if (!(!error.response)) {
              // A requisição foi feita e o servidor respondeu com um código de status
              // que sai do alcance de 2xx
              toast.error(error.response.status + " " + error.response.data.error);
            } else if (error.request) {
              // A requisição foi feita mas nenhuma resposta foi recebida
              // `error.request` é uma instância do XMLHttpRequest no navegador e uma instância de
              // http.ClientRequest no node.js
              toast.error(error.request);
            } else {
              // Alguma coisa acontenceu ao configurar a requisição que acionou este erro.
              toast.error('Error', error.message);
            }
            toast.error(error.config);
      });
    }catch( error : any){
      toast.error( error);
    }
  }

  const postTransferencia = async(e:React.FormEvent<HTMLFormElement>) =>{
    e.preventDefault();     
    try{
      const res = await axios.put(`http://localhost:8001/banco/conta/transferencia`,{
        "from": numContaDebito,
        "to": numContaCredito,
        "amount": valorTransferencia
      })
      .then((response)=>{toast.success(response.data.message)})
      .catch((error)=>{
          if (!(!error.response)) {
              // A requisição foi feita e o servidor respondeu com um código de status
              // que sai do alcance de 2xx
              toast.error(error.response.status + " " + error.response.data.error);
            } else if (error.request) {
              // A requisição foi feita mas nenhuma resposta foi recebida
              // `error.request` é uma instância do XMLHttpRequest no navegador e uma instância de
              // http.ClientRequest no node.js
              toast.error(error.request);
            } else {
              // Alguma coisa acontenceu ao configurar a requisição que acionou este erro.
              toast.error('Error', error.message);
            }
            toast.error(error.config);
      });
    }catch( error : any){
      toast.error( error);
    }
  }

  const postRender = async(e:React.FormEvent<HTMLFormElement>) =>{
    e.preventDefault();     
    try{
      const res = await axios.put(`http://localhost:8001/banco/conta/rendimento`,{"taxa": taxaRendimento})
      .then((response)=>{toast.success(response.data.message)})
      .catch((error)=>{
          if (!(!error.response)) {
              // A requisição foi feita e o servidor respondeu com um código de status
              // que sai do alcance de 2xx
              toast.error(error.response.status + " " + error.response.data.error);
            } else if (error.request) {
              // A requisição foi feita mas nenhuma resposta foi recebida
              // `error.request` é uma instância do XMLHttpRequest no navegador e uma instância de
              // http.ClientRequest no node.js
              toast.error(error.request);
            } else {
              // Alguma coisa acontenceu ao configurar a requisição que acionou este erro.
              toast.error('Error', error.message);
            }
            toast.error(error.config);
      });
    }catch( error : any){
      toast.error( error);
    }
  }

    return (
        <>
        <FormContainer onSubmit={getConta}>
            <InputArea>
                <Label>Numero da conta</Label>
                <Input name="numero" type= "number" onChange={(e)=>setContaNum(parseInt(e.target.value))}></Input>
            </InputArea>
            <Button type="submit">pesquisar</Button>
        </FormContainer>
        <Table>
            <Thead>
                <Tr>
                    <Th>Pontuacao</Th>
                    <Th>Saldo</Th>
                    <Th>Tipo</Th>
                </Tr>
                    <Tr>
                        <Td>{!conta ? "" : "" + conta.pontuacao + ""}</Td>
                        <Td>{!conta ? "" : "" + conta.saldo + ""}</Td>
                        <Td>{!conta ? "" : conta.tipo}</Td>
                    </Tr>
            </Thead>
        </Table>
        <h2>Fazer Credito</h2>
        <FormContainer onSubmit={postCredito}>
            <InputArea>
                <Label>Numero da conta</Label>
                <Input name="numero" type= "number" onChange={(e)=>setContaNum(parseInt(e.target.value))}></Input>
            
            </InputArea>
            <InputArea>
                <Label>Valor</Label>
                <Input name="valor" type= "number" onChange={(e)=>setValorCredito(parseInt(e.target.value))}></Input>
            
            </InputArea>
            <Button type="submit">Creditar</Button>
        </FormContainer>
        <h2>Fazer Debito</h2>
        <FormContainer onSubmit={postDebito}>
            <InputArea>
                <Label>Numero da conta</Label>
                <Input name="numero" type= "number" onChange={(e)=>setContaNum(parseInt(e.target.value))}></Input>
                
            </InputArea>
            <InputArea>
                <Label>Valor</Label>
                <Input name="valor" type= "number" onChange={(e)=>setValorDebito(parseInt(e.target.value))}></Input>
            
            </InputArea>
            <Button type="submit"> Debitar</Button>
        </FormContainer>
        <h2>Transferência</h2>
        <FormContainer onSubmit={postTransferencia}>
            <InputArea>
                <Label>Numero da conta de Credito</Label>
                <Input name="numero" type= "number" onChange={(e)=>setNumContaCredito(parseInt(e.target.value))}></Input>
            </InputArea>
            <InputArea>
                <Label>Numero da conta de Debito</Label>
                <Input name="numero" type= "number" onChange={(e)=>setNumContaDebito(parseInt(e.target.value))}></Input>
            </InputArea>
            <InputArea>
                <Label>Valor</Label>
                <Input name="valor" type= "number" onChange={(e)=>setValorTransferencia(parseInt(e.target.value))}></Input>
            
            </InputArea>
            <Button type="submit"> Transferir</Button>
        </FormContainer>
        <h2>Render</h2>
        <FormContainer onSubmit={postRender}>
            <InputArea>
                <Label>Taxa</Label>
                <Input name="numero" type= "number" onChange={(e)=>setTaxaRendimento(parseInt(e.target.value))}></Input>
            </InputArea>
            <Button type="submit"> Render</Button>
        </FormContainer>
        </>
    );
};

export default Grid;
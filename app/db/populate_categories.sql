INSERT INTO expense_categories (name)
VALUES
('Habitação'),
('Transporte'),
('Saúde'),
('Educação'),
('Alimentação'),
('Vestuário'),
('Cuidado Pessoal'),
('Lazer'),
('Dívida'),
('Doação'),
('Investimento'),
('Depósito de Segurança'),
('Outro');

INSERT INTO expense_sub_categories (category_id, name)
VALUES
-- HABITAÇÃO
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Aluguel'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Condomínio'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Prestação'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Diarista'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Secretária'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Luz'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Telefone Fixo'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Gás'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Internet'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'IPTU'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Compra de Utensílio'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Supermercado (Limpeza)'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Seguro'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Conserto'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Piscineiro'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Jardineiro'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'TV Assinatura'),
((SELECT id FROM expense_categories WHERE name = 'Habitação'), 'Celular'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Outro');

-- TRANSPORTE
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Prestação de Veículo'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Seguro'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Multa'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Licenciamento'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Combustível'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Estacionamento'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'IPVA'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Troca de Óleo'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Troca de Pneu'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Lavagem de Veículo'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Manutenção de Veículo'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Transporte Público'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Uber/99/Táxi'),
((SELECT id FROM expense_categories WHERE name = 'Transporte'), 'Outro');

-- SAÚDE
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Plano de Saúde'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Seguro de Vida'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Medicamento'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Médico'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Dentista'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Academia'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Psicólogo'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Fisioterapia'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Óculos/Lente'),
((SELECT id FROM expense_categories WHERE name = 'Saúde'), 'Outro');

-- EDUCAÇÃO
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Mensalidade'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Material Escolar'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Fardamento'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Curso'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Faculdade'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Livro'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Transporte Escolar'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Passeio Escolar'),
((SELECT id FROM expense_categories WHERE name = 'Educação'), 'Outro');

-- ALIMENTAÇÃO
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Supermercado'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Feira'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Açougue'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Restaurante'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Lanchonete'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Padaria'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Delivery'),
((SELECT id FROM expense_categories WHERE name = 'Alimentação'), 'Outro');

-- VESTUÁRIO
((SELECT id FROM expense_categories WHERE name = 'Vestuário'), 'Roupa'),
((SELECT id FROM expense_categories WHERE name = 'Vestuário'), 'Calçado'),
((SELECT id FROM expense_categories WHERE name = 'Vestuário'), 'Acessório'),
((SELECT id FROM expense_categories WHERE name = 'Vestuário'), 'Outro');

-- CUIDADOS PESSOAIS
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Supermercado (Higiene Pessoal)'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Cabeleireiro'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Barbeiro'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Manicure'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Pedicure'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Estética'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Cosmético'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Depilação'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Massagem'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Spa'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Produto de Beleza'),
((SELECT id FROM expense_categories WHERE name = 'Cuidado Pessoal'), 'Outro');

-- LAZER
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Cinema'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Teatro'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Show'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Balada'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Viagem'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Passeio'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Evento'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Hobby'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Jogo'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Atividade Recreativa'),
((SELECT id FROM expense_categories WHERE name = 'Lazer'), 'Outro');

-- DÍVIDAS
((SELECT id FROM expense_categories WHERE name = 'Dívida'), 'Cheque'),
((SELECT id FROM expense_categories WHERE name = 'Dívida'), 'Empréstimo'),
((SELECT id FROM expense_categories WHERE name = 'Dívida'), 'Financiamento'),
((SELECT id FROM expense_categories WHERE name = 'Dívida'), 'Outro');

-- DOAÇÃO
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Instituição de Caridade'),
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Dízimo'),
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Oferta'),
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Família'),
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Apoio a Amigo/Parente'),
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Campanha Social'),
((SELECT id FROM expense_categories WHERE name = 'Doação'), 'Outro');

-- INVESTIMENTO
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Ações'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Fundo de Investimento'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Poupança'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Criptomoeda'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Imóvel'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Compra de Ativos'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Previdência'),
((SELECT id FROM expense_categories WHERE name = 'Investimento'), 'Outro');

-- DEPÓSITO DE SEGURANÇA
((SELECT id FROM expense_categories WHERE name = 'Depósito de Segurança'), 'Aluguel'),
((SELECT id FROM expense_categories WHERE name = 'Depósito de Segurança'), 'Contrato'),
((SELECT id FROM expense_categories WHERE name = 'Depósito de Segurança'), 'Outro');

-- OUTRO
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Taxa Conta Corrente'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Taxa Limite Cheque Especial'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Anuidade Cartão'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Animal Doméstico'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Plano Funerário'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Presentes'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Taxas'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Mesadas'),
((SELECT id FROM expense_categories WHERE name = 'Outro'), 'Pensão Alimentícia');


--------------------------------------------------------------------------------------------------------------------------

INSERT INTO income_categories (name)
VALUES
('Salário'),
('Pró Labore'),
('Aluguel Recebido'),
('Freelancer'),
('Pensão'),
('Empréstimo Recebido'),
('Mesada'),
('Seguro'),
('Comissão'),
('Bônus'),
('Herança'),
('13º Salário'),
('Rendimento de Investimento'),
('Venda de Bens'),
('Venda de Ações'),
('Venda de Imóvel'),
('Venda de Veículo'),
('Doação Recebida'),
('Outro');
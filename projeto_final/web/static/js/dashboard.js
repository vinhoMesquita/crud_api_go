// Obter a data atual no formato "YYYY-MM-DD"
const hoje = new Date().toISOString().split('T')[0];

// DefiniNdo o id so usuario que sera utilizado na página
const userId = 1;

document.addEventListener('DOMContentLoaded', function () {
    tippy('[data-tippy-content]', {
        placement: 'right', // Coloca o tooltip à direita do ícone
        arrow: true, // Adiciona uma seta ao tooltip
    });
});

// Encontrar o elemento de entrada da data pelo ID
const inputDate = document.getElementById('data');
const startDateInput = document.getElementById('start-date');
const endDateInput = document.getElementById('end-date');

// Definir a data máxima como a data atual para que nao possa ser ccolocodas datas no futuro
inputDate.max = hoje;
startDateInput.max = hoje;
endDateInput.max = hoje;

// Função para verificar se já existe uma entrada para a data e o ID do usuário
function verificarEntradaExistente(date) {
    // Constrói a requisição GET com os parâmetros necessários para verificar se já existe uma entrada nessa data com o id do usuario
    return fetch(`/user/frequencia/?id_user=${userId}&start_date=${date}&end_date=${date}`)
        // Inicia a requisição fetch
        .then(response => {
            // Verifica se a resposta da requisição é bem-sucedida
            if (response.ok) {
                //converte a resposta para JSON e retorna essa promise
                return response.json();
            } else {
                // lanca um erro se der algo errado
                throw new Error('Erro ao verificar entrada existente');
            }
        });
}

// Função para atualizar a entrada existente
function updateFrequencia(date) {
    const t1 = Number(document.getElementById('t1').value);
    const t2 = Number(document.getElementById('t2').value);
    const t3 = Number(document.getElementById('t3').value);


    // Verificando se entrada e maior que zero e inteiro
    if (Number.isInteger(t1) && t1 >= 0 && Number.isInteger(t2) && t2 >= 0 && Number.isInteger(t3) && t3 >= 0) {
        // chamada do endpoint PATCH para atualizar valores na tabela de frequencia
        fetch(`/user/frequencia/update/`, {
            method: 'PATCH',
            headers: {
                "Content-Type": "application/json"
            },
            // Criando o body do json para atualizar os dados
            body: JSON.stringify({ "turno manha": t1, "turno tarde": t2, "turno noite": t3, "id_user": userId, "data": date })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Erro ao atualizar entrada existente');
            }
            return response.json();
        })
        .then(data => {
            console.log(data);
            Swal.fire(
                "Atualizado!",
                "A entrada foi atualizada com sucesso, grafico atualizado para os ultimos 30 dias",
                "success"
            );

            // redefinindo a o grafico para ultimos 30 dias
            const endDate = new Date().toISOString().split("T")[0];
            const startDate = new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).toISOString().split("T")[0];
            updateBarChart(startDate, endDate);
            updateLineChart(startDate, endDate);
        })
        .catch(error => {
            console.error("Erro ao atualizar a entrada de frequência:", error);
        });
    }
}

// Função para criar uma nova entrada de frequência
function createFrequencia(event) {
    event.preventDefault(); // Impede que a pagina recarregue
    const t1 = Number(document.getElementById("t1").value);
    const t2 = Number(document.getElementById("t2").value);
    const t3 = Number(document.getElementById("t3").value);
    const date = document.getElementById("data").value;

    // Verificar se já existe uma entrada para a data e o ID do usuário
    verificarEntradaExistente(date)
        .then(data => {
            if (data && data.length > 0) {
                // Caso verificarEntrada retorne com valores
                Swal.fire({
                    title: "Entrada já existente",
                    text: "Já existe uma entrada para esta data e ID de usuário. Deseja atualizar os dados existentes?",
                    icon: "warning",
                    showCancelButton: true,
                    confirmButtonText: "Sim, atualizar",
                    cancelButtonText: "Cancelar"
                }).then(result => {
                    if (result.isConfirmed) {
                        updateFrequencia(date);
                    }
                });
            } else {
                // verificacao de valores passados
                if (Number.isInteger(t1) && t1 >= 0 && Number.isInteger(t2) && t2 >= 0 && Number.isInteger(t3) && t3 >= 0) {
                    // Chamada do endpoint Post de CreateFrequencia
                    fetch(`/user/frequencia/`, {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json"
                        },
                        body: JSON.stringify({ "turno manha": t1, "turno tarde": t2, "turno noite": t3, "id_user": userId, "data": date })
                    })
                    .then(response => {
                        if (!response.ok) {
                            Swal.fira(
                                "Error!",
                                "A criacao da nova frequencia falhou"
                            );
                            throw new Error('Erro ao criar nova entrada de frequência');
                        }
                        return response.json();
                    })
                    .then(data => {
                        // se tudo ocorrer como o planejado
                        console.log(data);
                        Swal.fire(
                            "Adicionado!",
                            "A entrada foi feita com sucesso, grafico atualizado para os ultimos 30 dias",
                            "sucesso"
                        );

                        const endDate = hoje
                        const startDate = new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).toISOString().split('T')[0];
                        updateBarChart(startDate, endDate);
                        updateLineChart(startDate, endDate);
                    })
                    .catch(error => console.error('Erro:', error));
                }
            }
        })
        .catch(error => console.error('Erro ao verificar entrada existente e entrada de dados:', error));
}

// Função para buscar e atualizar o gráfico de barras
function updateBarChart(startDate, endDate) {
    let url = `/user/frequencia/?id_user=${userId}`;

    // Verifica se startDate foi definido e adiciona à URL
    if (startDate) {
        url += `&start_date=${startDate}`;
    }

    // Verifica se endDate foi definido e adiciona à URL
    if (endDate) {
        url += `&end_date=${endDate}`;
    }

    // Endpoint de GetByFilters para preencher o gráfico
    fetch(url)// Fucao executar o endpoint que retorna com as freqeuncia correspondentes
        .then(response => response.json())// Reposta do fetch convertida em um objeto JSON
        .then(data => { // Resultado do JSON
            if (!data || data.length === 0) {
                Swal.fire(
                    "Sem Dados",
                    "Nenhuma entrada de frequência encontrada para o período selecionado",
                    "warning"
                );
                return;
            }
            // Criando variaveis para acumular as frequencias de cada turno
            let totalT1 = 0;
            let totalT2 = 0;
            let totalT3 = 0;

            data.forEach(frequencia => {
                totalT1 += frequencia['turno manha'] || 0; // Acumulando os valores em cada uma das variaveis
                totalT2 += frequencia['turno tarde'] || 0;
                totalT3 += frequencia['turno noite'] || 0;
            });

            const dataTable = google.visualization.arrayToDataTable([ // Configuracoes do grafico utilizado
                ['Turno', 'Total'],
                ['Manhã', totalT1],
                ['Tarde', totalT2],
                ['Noite', totalT3]
            ]);

            const options = { // Opocoes que compoem o grafico gerado como titulo e cores
                title: 'Frequência por Turno',
                legend: { position: 'top' },
                colors: ['#FF6384', '#36A2EB', '#FFCE56'],
                animation: {
                    startup: true,
                    duration: 1000,
                    easing: 'out',
                }
            };

            // Intanciação do grafico utilizando as configurações criadas anteriormente
            const chart = new google.visualization.ColumnChart(document.getElementById('bar-chart'));
            chart.draw(dataTable, options);
        })
        .catch(error => console.error('Erro:', error));
}

// Função para buscar e atualizar o gráfico de linha
function updateLineChart(startDate, endDate) {
    let url = `/user/frequencia/?id_user=${userId}`;

    // Verifica se startDate foi definido e adiciona à URL
    if (startDate) {
        url += `&start_date=${startDate}`;
    }

    // Verifica se endDate foi definido e adiciona à URL
    if (endDate) {
        url += `&end_date=${endDate}`;
    }

    // Endpoint de GetByFilters para criação do gráfico de linhas
    fetch(url)
        .then(response => response.json())
        .then(data => {
            // Criacao da variavel frequencia semanal
            // Servira para guarda o valores Chave e valor para construcao do grafico
            const weeklyFrequency = {};

            // Itera sobre os dados recebidos para calcular a frequência semanal
            data.forEach(frequencia => {
                const day = frequencia['dia_semana']; // Criando um dia com base no conteudo do dia_semana do JSON
                if (!weeklyFrequency[day]) {
                    weeklyFrequency[day] = 0;
                }
                if (frequencia['turno manha'] !== undefined && frequencia['turno manha'] !== null) {
                    weeklyFrequency[day] += frequencia['turno manha'];
                }
                if (frequencia['turno tarde'] !== undefined && frequencia['turno tarde'] !== null) {
                    weeklyFrequency[day] += frequencia['turno tarde'];
                }
                if (frequencia['turno noite'] !== undefined && frequencia['turno noite'] !== null) {
                    weeklyFrequency[day] += frequencia['turno noite'];
                }
            });

            const dataTable = new google.visualization.DataTable();
            dataTable.addColumn('string', 'Dia da Semana');
            dataTable.addColumn('number', 'Total de Alunos');

            const diasDaSemana = {
                'Sunday': 'Domingo',
                'Monday': 'Segunda-feira',
                'Tuesday': 'Terça-feira',
                'Wednesday': 'Quarta-feira',
                'Thursday': 'Quinta-feira',
                'Friday': 'Sexta-feira',
                'Saturday': 'Sábado'
            };

            Object.entries(weeklyFrequency).forEach(([day, total]) => {
                dataTable.addRow([diasDaSemana[day], total]);
            });

            const options = {
                title: 'Frequência por Dia da Semana',
                legend: { position: 'top' },
                colors: ['#36A2EB'],
                animation: {
                    startup: true,
                    duration: 1000,
                    easing: 'out',
                }
            };

            const chart = new google.visualization.LineChart(document.getElementById('line-chart'));
            chart.draw(dataTable, options);
        })
        .catch(error => console.error('Erro:', error));
}

// Função para buscar as informações do usuário
function getUserInfo() {
    fetch(`/user/?userid=${userId}`)
        .then(response => response.json())
        .then(data => {
            if (data) {
                document.getElementById("userid").textContent = data.id;
                document.getElementById("username").textContent = data.nome;
            } else {
                Swal.fire(
                    "Usuario inexistente",
                    "Nenhum usuário exitente com esse ID",
                    "warning"
                );
            }
        })
        .catch(error => console.error('Erro:', error));
}

getUserInfo(userId);

// Carregar o Google Charts
google.charts.load('current', { packages: ['corechart'] });
// Quando o Google Charts for carregado, execute a função de callback
google.charts.setOnLoadCallback(() => {

    // Definir a data de início e a data de fim para os últimos 30 dias
    const endDate = hoje;
    const startDate = new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).toISOString().split('T')[0];

    // Atualizar os gráficos com a data de início e a data de fim definidas
    updateBarChart(startDate, endDate);
    updateLineChart(startDate, endDate);

    // Função para atualizar gráficos ao mudar os inputs de data
    const updateChartsOnDateChange = () => {
        const startDateValue = document.getElementById("start-date").value;
        const endDateValue = document.getElementById("end-date").value;

        // Verifica se a data de início não é maior que a data de fim
        if (startDateValue && endDateValue && new Date(startDateValue) > new Date(endDateValue)) {
            Swal.fire(
                "Data Inválida",
                "A data de início não pode ser maior que a data de fim.",
                "error"
            );
            return;
        }

        // Atualizar os gráficos com os novos valores de data
        updateBarChart(startDateValue || undefined, endDateValue || undefined);
        updateLineChart(startDateValue || undefined, endDateValue || undefined);
    };

    // Adicionar eventos de input nos campos de data para atualizar os gráficos automaticamente
    document.getElementById("start-date").addEventListener("input", updateChartsOnDateChange);
    document.getElementById("end-date").addEventListener("input", updateChartsOnDateChange);
});

// Adicionar nova entrada
document.getElementById("new-entry-form").addEventListener("submit", createFrequencia);

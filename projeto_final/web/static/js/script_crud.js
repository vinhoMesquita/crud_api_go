



document.addEventListener('DOMContentLoaded', function () {
    fetchAlunos(); // Chama a função para buscar e exibir os alunos ao carregar a página
});


// Função para buscar todos os alunos
function fetchAlunos() {
    const url = 'http://localhost:5501/alunos'; // URL da sua API

    fetch(url)
        // console.log(aluno)
        .then(response => response.json())
        .then(aluno => {
            displayAlunos(aluno);
        })
        .catch(error => {
            console.error('Erro ao buscar alunos:', error);
            showMessage('Erro ao buscar alunos!', 'error');
        });
}

// Função para buscar aluno por ID
function fetchAlunoById(id) {
    const url = `http://localhost:5501/alunos/${id}`; // URL da sua API para buscar por ID

    fetch(url)
        .then(response => response.json())
        .then(aluno => {
            displayAluno(aluno);
        })
        .catch(error => {
            console.error(`Erro ao buscar aluno com ID ${id}:`, error);
            showMessage(`Erro ao buscar aluno com ID ${id}!`, 'error');
        });
}

// Função para cadastrar um novo aluno
function cadastrarAluno(aluno) {
    const url = 'http://localhost:5501/alunos'; // URL da sua API para cadastrar

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(aluno),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Erro ao cadastrar aluno: ${response.statusText}`);
        }
        showMessage('Aluno cadastrado com sucesso!', 'success');
        fetchAlunos(); // Atualiza a tabela após o cadastro
    })
    .catch(error => {
        console.error('Erro ao cadastrar aluno:', error);
        showMessage('Erro ao cadastrar aluno!', 'error');
    });
}

// Função para atualizar aluno por ID
function atualizarAluno(aluno) {
    const url = `http://localhost:5501/alunos/${aluno.id_aluno}`; // URL da sua API para atualizar

    fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(aluno),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Erro ao atualizar aluno: ${response.statusText}`);
        }
        showMessage('Aluno atualizado com sucesso!', 'success');
        fetchAlunos(); // Atualiza a tabela após a atualização
    })
    .catch(error => {
        console.error('Erro ao atualizar aluno:', error);
        showMessage('Erro ao atualizar aluno!', 'error');
    });
}

// Função para deletar aluno por ID
function deletarAluno(id) {
    const url = `http://localhost:5501/alunos/${id}`; // URL da sua API para deletar

    fetch(url, {
        method: 'DELETE',
    })
    .then(response => {
        if (!response.ok) {
            throw new Error(`Erro ao deletar aluno: ${response.statusText}`);
        }
        showMessage('Aluno deletado com sucesso!', 'success');
        fetchAlunos(); // Atualiza a tabela após a exclusão
    })
    .catch(error => {
        console.error(`Erro ao deletar aluno com ID ${id}:`, error);
        showMessage(`Erro ao deletar aluno com ID ${id}!`, 'error');
    });
}

// Função para exibir todos os alunos na tabela
function displayAlunos(alunos) {
    const tableBody = document.querySelector('#alunos-table tbody');
    tableBody.innerHTML = ''; // Limpa a tabela antes de adicionar novos dados

    alunos.forEach(aluno => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${aluno.id_aluno}</td>
            <td>${aluno.name_aluno}</td>
            <td>${aluno.age_aluno}</td>
            <td>${aluno.bf_aluno}</td>
            <td>${aluno.mm_aluno}</td>
            <td>${aluno.altura_aluno}</td>
            <td>${aluno.peso_aluno}</td>
            <td>${aluno.imc_aluno}</td>
        `;
        tableBody.appendChild(row);
    });
}

// Função para exibir os dados de um único aluno em um formulário de edição
function displayAluno(aluno) {
    document.getElementById('id-editar').value = aluno.id_aluno;
    document.getElementById('nome-editar').value = aluno.name_aluno;
    document.getElementById('idade-editar').value = aluno.age_aluno;
    document.getElementById('bf-editar').value = aluno.bf_aluno;
    document.getElementById('massa-editar').value = aluno.mm_aluno;
    document.getElementById('altura-editar').value = aluno.altura_aluno;
    document.getElementById('peso-editar').value = aluno.peso_aluno;
}

// Função para exibir mensagens na tela
function showMessage(message, type) {
    const messageBox = document.getElementById('message');
    messageBox.textContent = message;
    messageBox.className = type; // Define a classe CSS para o estilo de mensagem de acordo com o tipo
    messageBox.style.display = 'block';

    setTimeout(() => {
        messageBox.style.display = 'none'; // Oculta a mensagem após 3 segundos
    }, 3000);
}

// Event listener para o formulário de cadastro de aluno
document.getElementById('form-aluno').addEventListener('submit', function (event) {
    event.preventDefault();

    const aluno = {
        name_aluno: document.getElementById('nome').value,
        age_aluno: parseInt(document.getElementById('idade').value, 10),
        bf_aluno: parseFloat(document.getElementById('bf').value),
        mm_aluno: parseFloat(document.getElementById('massa').value),
        altura_aluno: parseFloat(document.getElementById('altura').value),
        peso_aluno: parseFloat(document.getElementById('peso').value)
    };

    if (!validateForm(aluno)) {
        showMessage('Preencha todos os campos!', 'error');
        return;
    }

    cadastrarAluno(aluno);
    document.getElementById('form-aluno').reset();
});

// Event listener para o formulário de busca de aluno por ID
document.getElementById('form-buscar').addEventListener('submit', function (event) {
    event.preventDefault();

    const id = parseInt(document.getElementById('id-buscar').value, 10);
    fetchAlunoById(id);
});

// Event listener para o formulário de deleção de aluno por ID
document.getElementById('form-deletar').addEventListener('submit', function (event) {
    event.preventDefault();

    const id = parseInt(document.getElementById('id-deletar').value, 10);
    deletarAluno(id);
});

// Event listener para o formulário de edição de aluno
document.getElementById('form-editar').addEventListener('submit', function (event) {
    event.preventDefault();

    const aluno = {
        id_aluno: parseInt(document.getElementById('id-editar').value, 10),
        name_aluno: document.getElementById('nome-editar').value,
        age_aluno: parseInt(document.getElementById('idade-editar').value, 10),
        bf_aluno: parseFloat(document.getElementById('bf-editar').value),
        mm_aluno: parseFloat(document.getElementById('massa-editar').value),
        altura_aluno: parseFloat(document.getElementById('altura-editar').value),
        peso_aluno: parseFloat(document.getElementById('peso-editar').value)
    };

    if (!validateForm(aluno)) {
        showMessage('Preencha todos os campos!', 'error');
        return;
    }

    atualizarAluno(aluno);
});

// Função para validar o formulário de aluno
function validateForm(aluno) {
    return aluno.name_aluno && aluno.age_aluno && aluno.bf_aluno && aluno.mm_aluno && aluno.altura_aluno && aluno.peso_aluno;
}

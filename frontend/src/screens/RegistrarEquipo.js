import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, Button, StyleSheet, ScrollView } from 'react-native';
import { Picker } from "@react-native-picker/picker";

const RegistrarEquipo = () => {
    const [teamName, setTeamName] = useState('');
    const [description, setDescription] = useState('');
    const [participants, setParticipants] = useState([{ name: '', email: '', phone: '' }]);
    const [robots, setRobots] = useState([{ name: '', description: '', categoryID: '' }]);
    const [categories, setCategories] = useState([]);
    const [errors, setErrors] = useState({});

    useEffect(() => {
        const hardcodedCategories = [
            { ID: 1, Nombre: 'Sumo', Descripcion: 'Robots de sumo' },
            { ID: 2, Nombre: 'Peleas', Descripcion: 'Robots de peleas de robots' },
            { ID: 3, Nombre: 'Seguidor de linea', Descripcion: 'Seguidores de linea' },
            { ID: 4, Nombre: 'Sumo', Descripcion: 'Robots de sumo' },
            { ID: 5, Nombre: 'Peleas', Descripcion: 'Robots de peleas de robots' },
            { ID: 6, Nombre: 'Seguidor de linea', Descripcion: 'Seguidores de linea' }
        ];
        setCategories(hardcodedCategories);
    }, []);

    const validate = () => {
        const newErrors = {};
        if (!teamName) newErrors.teamName = 'Team name is required';
        participants.forEach((participant, index) => {
            if (!participant.name) newErrors[`participantName${index}`] = 'Participant name is required';
            if (!participant.email) newErrors[`participantEmail${index}`] = 'Participant email is required';
            if (!participant.phone) newErrors[`participantPhone${index}`] = 'Participant phone is required';
        });
        robots.forEach((robot, index) => {
            if (!robot.name) newErrors[`robotName${index}`] = 'Robot name is required';
            if (!robot.categoryID) newErrors[`robotCategory${index}`] = 'Robot category is required';
        });
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleRegister = () => {
        if (!validate()) return;

        const teamData = {
            Nombre: teamName,
            Descripcion: description,
            Participantes: participants.map(p => ({
                Nombre: p.name,
                Correo: p.email,
                Telefono: p.phone
            })),
            Robots: robots.map(r => ({
                Nombre: r.name,
                Descripcion: r.description,
                CategoriaID: parseInt(r.categoryID)
            }))
        };
        console.log('Team Data:', teamData);
        // Aquí puedes hacer la llamada a la API para registrar el equipo
    };

    const handleParticipantChange = (index, field, value) => {
        const newParticipants = [...participants];
        newParticipants[index][field] = value;
        setParticipants(newParticipants);
    };

    const handleRobotChange = (index, field, value) => {
        const newRobots = [...robots];
        newRobots[index][field] = value;
        setRobots(newRobots);
    };

    const addParticipant = () => {
        setParticipants([...participants, { name: '', email: '', phone: '' }]);
    };

    const addRobot = () => {
        setRobots([...robots, { name: '', description: '', categoryID: '' }]);
    };

    const removeParticipant = (index) => {
        if (participants.length > 1) {
            const newParticipants = participants.filter((_, i) => i !== index);
            setParticipants(newParticipants);
        }
    };

    const removeRobot = (index) => {
        if (robots.length > 1) {
            const newRobots = robots.filter((_, i) => i !== index);
            setRobots(newRobots);
        }
    };

    return (
        <ScrollView style={styles.container}>
            <Text style={styles.title}>Register Team</Text>
            <TextInput
                style={styles.input}
                placeholder="Team Name"
                value={teamName}
                onChangeText={setTeamName}
            />
            {errors.teamName && <Text style={styles.error}>{errors.teamName}</Text>}
            <TextInput
                style={styles.input}
                placeholder="Description"
                value={description}
                onChangeText={setDescription}
            />
            <Text style={styles.subtitle}>Participants</Text>
            {participants.map((participant, index) => (
                <View key={index} style={styles.participantContainer}>
                    <TextInput
                        style={styles.input}
                        placeholder="Name"
                        value={participant.name}
                        onChangeText={value => handleParticipantChange(index, 'name', value)}
                    />
                    {errors[`participantName${index}`] && <Text style={styles.error}>{errors[`participantName${index}`]}</Text>}
                    <TextInput
                        style={styles.input}
                        placeholder="Email"
                        value={participant.email}
                        onChangeText={value => handleParticipantChange(index, 'email', value)}
                    />
                    {errors[`participantEmail${index}`] && <Text style={styles.error}>{errors[`participantEmail${index}`]}</Text>}
                    <TextInput
                        style={styles.input}
                        placeholder="Phone"
                        value={participant.phone}
                        onChangeText={value => handleParticipantChange(index, 'phone', value)}
                    />
                    {errors[`participantPhone${index}`] && <Text style={styles.error}>{errors[`participantPhone${index}`]}</Text>}
                    <Button title="Remove" onPress={() => removeParticipant(index)} />
                </View>
            ))}
            <Button title="Add Participant" onPress={addParticipant} />
            <Text style={styles.subtitle}>Robots</Text>
            {robots.map((robot, index) => (
                <View key={index} style={styles.robotContainer}>
                    <TextInput
                        style={styles.input}
                        placeholder="Name"
                        value={robot.name}
                        onChangeText={value => handleRobotChange(index, 'name', value)}
                    />
                    {errors[`robotName${index}`] && <Text style={styles.error}>{errors[`robotName${index}`]}</Text>}
                    <TextInput
                        style={styles.input}
                        placeholder="Description"
                        value={robot.description}
                        onChangeText={value => handleRobotChange(index, 'description', value)}
                    />
                    <Picker
                        selectedValue={robot.categoryID}
                        onValueChange={value => handleRobotChange(index, 'categoryID', value)}
                    >
                        {categories.map(category => (
                            <Picker.Item key={category.ID} label={category.Nombre} value={category.ID} />
                        ))}
                    </Picker>
                    {errors[`robotCategory${index}`] && <Text style={styles.error}>{errors[`robotCategory${index}`]}</Text>}
                    <Button title="Remove" onPress={() => removeRobot(index)} />
                </View>
            ))}
            <Button title="Add Robot" onPress={addRobot} />
            <Button title="Register" onPress={handleRegister} />
        </ScrollView>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 16,
        height: '85%',
    },
    title: {
        fontSize: 24,
        marginBottom: 16,
        textAlign: 'center',
    },
    subtitle: {
        fontSize: 20,
        marginTop: 16,
        marginBottom: 8,
    },
    input: {
        height: 40,
        borderColor: 'gray',
        borderWidth: 1,
        marginBottom: 12,
        paddingHorizontal: 8,
    },
    participantContainer: {
        marginBottom: 16,
    },
    robotContainer: {
        marginBottom: 16,
    },
    error: {
        color: 'red',
        marginBottom: 8,
    },
});

export default RegistrarEquipo;
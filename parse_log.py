import re
import matplotlib.pyplot as plt

# Function to parse the log file
def parse_log(log_file):
    lateral_3 = []
    lateral_2 = []
    lateral_1 = []
    epochs = []

    with open(log_file, 'r') as f:
        for line in f:
            match = re.search(r'\[lateral-3: (\d+\.\d+)\], \[lateral-2: (\d+\.\d+)\], \[lateral-1: (\d+\.\d+)\]', line)
            if match:
                lateral_3.append(float(match.group(1)))
                lateral_2.append(float(match.group(2)))
                lateral_1.append(float(match.group(3)))
                epoch_match = re.search(r'Epoch \[(\d+)/\d+\]', line)
                if epoch_match:
                    epochs.append(int(epoch_match.group(1)))

    return epochs, lateral_3, lateral_2, lateral_1

# Function to plot the chart
def plot_chart(epochs, lateral_3, lateral_2, lateral_1):
    # print(lateral_1)
    plt.figure(figsize=(10, 6))
    plt.plot(epochs, lateral_3, label='lateral-3')
    plt.plot(epochs, lateral_2, label='lateral-2')
    plt.plot(epochs, lateral_1, label='lateral-1')
    plt.title('Training Progress')
    plt.xlabel('Epochs')
    plt.ylabel('Laterals')
    plt.xticks(range(25))  # Adjust x-axis ticks
    plt.yticks([0, 1, 2])  # Adjust y-axis ticks
    plt.legend()
    plt.grid(True)
    plt.show()

# Parse the log file
log_file = './log/BGNet.txt'
epochs, lateral_3, lateral_2, lateral_1 = parse_log(log_file)

# Plot the chart
plot_chart(epochs, lateral_3, lateral_2, lateral_1)
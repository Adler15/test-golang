import paramiko
import time
import os
import sys


docker_file = sys.argv[1]
local_path = sys.argv[2]
docker_image = sys.argv[3]
host = sys.argv[4]
port = sys.argv[5]
username = sys.argv[6]
password = sys.argv[7]
target_path = sys.argv[8]



class SSHConnection:
    def __init__(self, host, port, username, local_path, docker_file, password, target_path,docker_image):
        self.host = host
        self.port = port
        self.docker_image = docker_image
        self.username = username
        self.password = password
        self.docker_file = docker_file
        self.target_path = target_path
        self.local_path = local_path
    
    def modify(self):
        if self.host == "192.168.130.109":
            os.popen("sed -i 's#bcf-basic-ms:0.1#%s#g' %s " % (self.docker_image, (self.local_path + self.docker_file)))
            time.sleep(1)
            os.popen("sed -i 's#192.168.130.109#%s#g' %s " % (self.host, (self.local_path + self.docker_file)))
        elif self.host == "192.168.130.117":
            os.popen("sed -i 's#bcf-basic-ms:0.1#%s#g' %s " % (self.docker_image, (self.local_path + self.docker_file)))
            time.sleep(1)
            os.popen("sed -i 's#192.168.130.109#%s#g' %s " % (self.host, (self.local_path + self.docker_file)))
            time.sleep(1)
            os.popen("sed -i 's#bcf-deploy-dist-v1.0.0#bcf-deploy-dist-v2.0.0#g' %s " % (self.local_path + self.docker_file))

    def run(self):
        self.modify()
        self.connect()
        self.cmd('rm -rf %s' % (self.target_path + self.docker_file))
        self.upload(self.local_path + self.docker_file, self.target_path + self.docker_file)
        self.cmd("docker rm -f bcf-basic-ms")
        time.sleep(2)
        self.cmd("docker-compose -f %s/%s -p k2data-bcf stop " % (self.target_path, self.docker_file))
        time.sleep(2)
        self.cmd("docker-compose -f %s/%s -p k2data-bcf rm -f " % (self.target_path, self.docker_file))
        time.sleep(2)
        self.cmd("docker-compose -f %s/%s -p k2data-bcf up -d " % (self.target_path, self.docker_file))

    def connect(self):
        transport = paramiko.Transport((self.host, self.port))
        transport.connect(username=self.username, password=self.password)
        self.__transport = transport

    def upload(self, local_path, target_path):
        sftp = paramiko.SFTPClient.from_transport(self.__transport)
        sftp.put(local_path, target_path)

    def cmd(self, command):
        ssh = paramiko.SSHClient()
        ssh._transport = self.__transport
        stdin, stdout, stderr = ssh.exec_command(command)
        result = stdout.read()
        print(result)

        return result



if __name__ == '__main__':

    obj = SSHConnection(host=host,
                        port=int(port),
                        username=username,
                        password=password,
                        target_path=target_path,
                        local_path=local_path,
                        docker_file=docker_file,
                        docker_image=docker_image)

    obj.run()
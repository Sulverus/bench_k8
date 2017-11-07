box.cfg{
    listen="3316";
    replication={"127.0.0.1:3315"}
}

function stop()
    box.cfg{replication={}}
    require('fiber').sleep(1.5)
end

function info()
    accounts = box.space.accounts:select{}
    sum = 0
    for _, acc in pairs(accounts) do
        sum = sum + acc[2]
    end
    return accounts, sum
end

function send(from, to, amount)
    box.begin()
    box.space.accounts:update({from}, {{"-", 2, amount}})
    box.space.accounts:update({to}, {{"+", 2, amount}})
    box.commit()
end

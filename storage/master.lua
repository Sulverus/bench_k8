box.cfg{
    listen="3315"
}

box.schema.create_space("accounts")
box.space.accounts:create_index("primary")

problem = false

function problem()
    problem = true
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
    if problem == true then
            require('fiber').sleep(1)
            return
    end
    box.begin()
    --require('log').info("from %d, to %d, x=%d", from, to, amount)
    box.space.accounts:update({from}, {{"-", 2, amount}})
    box.space.accounts:update({to}, {{"+", 2, amount}})
    box.commit()
end

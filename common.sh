export TMPDIR=${TMPDIR:-/tmp}

function setuprepo() {
  for row in $(jq -cr '. | select(.source.repos != null) | .source.repos[]' < $payload); do
    name=$(echo $row | jq -cr .name)
    url=$(echo $row | jq -cr .url)
    username=$(echo $row | jq -cr .username)
    password=$(echo $row | jq -cr .password)
    OPTS=""

    if [[ $username != "null" && $password != "null" ]]; then
      OPTS="--username $username --password $password"
    fi

    helm repo add $name $url $OPTS
  done
}